defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> parse()
    |> scale()
    |> simulate()
    |> Enum.filter(fn
      {_, ?[} ->
        true

      {_, _} ->
        false
    end)
    |> Enum.map(&elem(&1, 0))
    |> Enum.map(&gps/1)
    |> Enum.sum()
  end

  def simulate(%{map: map, moves: []}) do
    map
  end

  def simulate(%{map: map, robot_pos: r, moves: [m | tail]}) do
    {updated_map, updated_rpos} = apply_move(map, r, m)
    simulate(%{map: updated_map, robot_pos: updated_rpos, moves: tail})
  end

  def apply_move(map, r, move) do
    offset =
      case move do
        ?> -> {0, 1}
        ?< -> {0, -1}
        ?^ -> {-1, 0}
        ?v -> {1, 0}
      end

    target = add(r, offset)

    case(Map.get(map, target)) do
      # simple move to empty box
      ?. ->
        {map
         |> Map.put(r, ?.)
         |> Map.put(target, ?@), target}

      # wall tile, no move, map unchanged
      ?# ->
        {map, r}

      # move 1+ boxes
      box when box in [?[, ?]] and move in [?>, ?<] ->
        case move_boxes_xaxis(map, target, offset) do
          {:ok, new_map} ->
            {new_map
             |> Map.put(r, ?.)
             |> Map.put(target, ?@), target}

          {:error} ->
            {map, r}
        end

      box when box in [?[, ?]] and move in [?^, ?v] ->
        case move_boxes_yaxis(map, target, offset) do
          {:ok, new_map} ->
            {new_map
             |> Map.put(r, ?.)
             |> Map.put(target, ?@), target}

          {:error} ->
            {map, r}
        end
    end
  end

  def move_boxes_yaxis(map, target, offset) do
    to_shift =
      [target] ++
        case Map.get(map, target) do
          ?[ ->
            [add(target, {0, 1})]

          ?] ->
            [add(target, {0, -1})]
        end

    case shiftable_yaxis?(map, to_shift, offset, MapSet.new()) do
      {true, dedup} ->
        movable_boxes_yaxis =
          dedup
          |> Enum.filter(fn x ->
            z = Map.get(map, x)
            z == ?[ or z == ?]
          end)
          |> Enum.flat_map(fn {r, c} ->
            case Map.get(map, {r, c}) do
              ?] -> [{r, c - 1}, {r, c}]
              ?[ -> [{r, c}, {r, c + 1}]
            end
          end)
          |> Enum.uniq()

        {:ok, shift_yaxis(map, movable_boxes_yaxis, offset)}

      false ->
        {:error}
    end
  end

  defp shiftable_yaxis?(map, [], offset, dedup) do
    {true, dedup}
  end

  defp shiftable_yaxis?(map, [h | tail], offset, dedup) do
    next_row_pos = add(h, offset)

    if MapSet.member?(dedup, h) do
      shiftable_yaxis?(map, tail, offset, dedup)
    else
      dedup = MapSet.put(dedup, h)

      case Map.get(map, next_row_pos) do
        ?# ->
          false

        ?. ->
          shiftable_yaxis?(map, tail, offset, dedup)

        ?[ ->
          tail = [next_row_pos | tail]
          tail = [add(next_row_pos, {0, 1}) | tail]

          shiftable_yaxis?(map, tail, offset, dedup)

        ?] ->
          tail = [next_row_pos | tail]
          tail = [add(next_row_pos, {0, -1}) | tail]

          shiftable_yaxis?(map, tail, offset, dedup)
      end
    end
  end

  # going up
  defp shift_yaxis(map, movable_boxes_yaxis, {-1, 0}) do
    movable_boxes_yaxis
    |> Enum.sort(:asc)
    |> Enum.chunk_every(2)
    |> Enum.reduce(map, fn [{lr, lc}, {rr, rc}], acc ->
      %{acc | {lr - 1, lc} => ?[, {rr - 1, rc} => ?], {lr, lc} => ?., {rr, rc} => ?.}
    end)
  end

  # going down
  defp shift_yaxis(map, movable_boxes_yaxis, {1, 0}) do
    movable_boxes_yaxis
    |> Enum.sort(:desc)
    |> Enum.chunk_every(2)
    # lX and rX swapped on purpose - the order would be wrong cause of sort :desc
    |> Enum.reduce(map, fn [{rr, rc}, {lr, lc}], acc ->
      %{acc | {lr + 1, lc} => ?[, {rr + 1, rc} => ?], {lr, lc} => ?., {rr, rc} => ?.}
    end)
  end

  def move_boxes_xaxis(map, first_box_pos, offset) do
    movable_boxes = movable_boxes_xaxis(map, first_box_pos, offset, 1)

    if movable_boxes == 0 do
      {:error}
    else
      {rows, cols} = offset

      upd_map =
        shift_xaxis(
          map,
          first_box_pos,
          add(first_box_pos, {rows * movable_boxes, cols * movable_boxes}),
          offset,
          Map.get(map, first_box_pos)
        )

      {:ok, upd_map}
    end
  end

  def shift_xaxis(map, cur, fin, offset, buffer) do
    if cur == fin do
      map
    else
      next = Map.get(map, add(cur, offset))
      new_map = Map.put(map, add(cur, offset), buffer)
      shift_xaxis(new_map, add(cur, offset), fin, offset, next)
    end
  end

  def movable_boxes_xaxis(map, box_pos, offset, n) do
    case Map.get(map, add(box_pos, offset)) do
      ?. -> n
      ?# -> 0
      ?] -> movable_boxes_xaxis(map, add(box_pos, offset), offset, n + 1)
      ?[ -> movable_boxes_xaxis(map, add(box_pos, offset), offset, n + 1)
    end
  end

  def add({x, y}, {ox, oy}) do
    {x + ox, y + oy}
  end

  def gps({row, col}) do
    100 * row + col
  end

  def parse(input) do
    [map, moves] =
      input
      |> String.split("\n\n", trim: true)

    warehouse_map =
      map
      |> String.to_charlist()
      |> Enum.reduce(
        {0, 0, %{}},
        fn
          ?\n, {row, _, acc} ->
            {row + 1, 0, acc}

          char, {row, col, acc} ->
            {
              row,
              col + 1,
              Map.put_new(acc, {row, col}, char)
            }
        end
      )
      |> elem(2)

    robot_pos = Enum.find(warehouse_map, fn {_k, v} -> v == ?@ end) |> elem(0)

    %{
      robot_pos: robot_pos,
      map: warehouse_map,
      moves: moves |> String.to_charlist() |> Enum.filter(fn x -> x != ?\n end)
    }
  end

  def scale(%{
        robot_pos: _,
        map: map,
        moves: moves
      }) do
    new_map =
      Enum.reduce(map, %{}, fn
        {{row, col}, ?@}, acc ->
          Map.put(acc, {row, col * 2}, ?@)
          |> Map.put({row, col * 2 + 1}, ?.)

        {{row, col}, ?O}, acc ->
          Map.put(acc, {row, col * 2}, ?[)
          |> Map.put({row, col * 2 + 1}, ?])

        {{row, col}, x}, acc ->
          Map.put(acc, {row, col * 2}, x)
          |> Map.put({row, col * 2 + 1}, x)
      end)

    {nrp, _} =
      new_map
      |> Enum.find(fn {_, x} -> x == ?@ end)

    %{
      moves: moves,
      map: new_map,
      robot_pos: nrp
    }
  end

  def pretty_print(map) do
    IO.puts("\n\n")

    Map.keys(map)
    |> Enum.group_by(fn {row, _col} -> row end)
    |> Enum.map(fn {_k, v} ->
      v
      |> Enum.sort()
      |> Enum.map(fn x -> Map.get(map, x) end)
    end)
    |> Enum.each(fn row -> IO.puts(row) end)

    IO.puts("\n\n")
  end
end
