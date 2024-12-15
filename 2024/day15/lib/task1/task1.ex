defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> parse()
    |> simulate()
    |> Enum.filter(fn
      {_, ?O} ->
        true

      {_, _} ->
        false
    end)
    |> Enum.map(&elem(&1, 0))
    |> Enum.map(&gps/1)
    |> Enum.sum()
  end

  defp simulate(%{map: map, moves: []}) do
    map
  end

  defp simulate(%{map: map, robot_pos: r, moves: [m | tail]}) do
    {updated_map, updated_rpos} = apply_move(map, r, m)
    simulate(%{map: updated_map, robot_pos: updated_rpos, moves: tail})
  end

  defp apply_move(map, r, move) do
    offset =
      case move do
        ?^ -> {-1, 0}
        ?> -> {0, 1}
        ?< -> {0, -1}
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
      ?O ->
        case move_boxes(map, target, offset) do
          {:ok, new_map} ->
            {new_map
             |> Map.put(r, ?.)
             |> Map.put(target, ?@), target}

          {:error} ->
            {map, r}
        end
    end
  end

  def move_boxes(map, first_box_pos, offset) do
    movable_boxes = movable_boxes(map, first_box_pos, offset, 1)

    if movable_boxes == 0 do
      {:error}
    else
      {rows, cols} = offset

      {:ok,
       map
       |> Map.put(first_box_pos, ?.)
       |> Map.put(add(first_box_pos, {rows * movable_boxes, cols * movable_boxes}), ?O)}
    end
  end

  defp movable_boxes(map, box_pos, offset, n) do
    case Map.get(map, add(box_pos, offset)) do
      ?. -> n
      ?# -> 0
      ?O -> movable_boxes(map, add(box_pos, offset), offset, n + 1)
    end
  end

  defp add({x, y}, {ox, oy}) do
    {x + ox, y + oy}
  end

  defp gps({row, col}) do
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
