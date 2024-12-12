defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {rows, cols, parsed} = parse(input)

    parsed
    |> Map.keys()
    |> Enum.flat_map(fn k ->
      build_regions({rows, cols}, Map.get(parsed, k), [])
      |> Enum.map(fn region ->
        region # |> IO.inspect(label: :region)
        |> cost_of_fence({rows, cols})
      end)
    end)
    |> Enum.sum()
  end

  def build_regions(size, positions, regions) do
    if Enum.empty?(positions) do
      regions
    else
      first_pos = positions |> MapSet.to_list() |> Enum.at(0)
      region = build_current_region(size, first_pos, positions |> MapSet.delete(first_pos) |> MapSet.to_list(), [first_pos])
      new_positions = MapSet.difference(positions, MapSet.new(region))
      build_regions(size, new_positions, [region | regions])
    end

    # 1 take first position
    # 2 generate all potential neighboors (not outside of map bounds)
    # 3 check if any neighboors are part of positions list
    # 4 if yes - add them to current region started at first pos. repeat 2-4 for each of them

    # 5 remove all positions existing in current region from positions
    # 6 repeat 1-5
  end

  def build_current_region(_, pos, [], []) do
    [pos]
  end

  def build_current_region(_, _, [], region) do
    region
  end

  def build_current_region(size, pos, positions, region) do
    # IO.inspect(pos, label: :pos)
    # IO.inspect(positions, label: :positions)
    # IO.inspect(region, label: :region)

    neighboors = valid_neighboors(size, pos)
    # IO.inspect(neighboors, label: :neighboors)
    existing_neighboors = Enum.filter(positions, fn x -> Enum.member?(neighboors, x) end)
    # IO.inspect(existing_neighboors, label: :existing_neighboors)

    if Enum.empty?(existing_neighboors) do
      region
    else
      new_region = existing_neighboors ++ region
      # IO.inspect(new_region, label: :new_region)
      new_positions = Enum.filter(positions, fn x -> not Enum.member?(neighboors, x) end)

      Enum.flat_map(existing_neighboors, fn x ->
        build_current_region(size, x, new_positions, new_region)
      end)
      |> Enum.uniq()
    end
  end

  def valid_neighboors({mrow, mcol}, {x, y}) do
    [
      {x + 1, y},
      {x - 1, y},
      {x, y + 1},
      {x, y - 1}
    ]
    |> Enum.filter(fn {row, col} ->
      row >= 0 and col >= 0 and row <= mrow and col <= mcol
    end)
  end

  def cost_of_fence(region, size) do
    area = Enum.count(region)
    perimeter = 4 * area - direct_neighboors(size, region, 0) * 2

    area * perimeter
  end

  defp direct_neighboors(_, [], acc) do
    acc
  end

  defp direct_neighboors(size, [pos| region], acc) do
    all_neighboors = valid_neighboors(size, pos)
    new_acc = acc + length(Enum.filter(region, fn x -> Enum.member?(all_neighboors, x) end))
    direct_neighboors(size, region, new_acc)
  end

  @spec parse(binary()) :: {any(), number(), map()}
  def parse(input) do
    {rows, cols, map} =
      input
      |> String.trim()
      |> String.to_charlist()
      |> Enum.reduce(
        {0, 0, %{}},
        fn
          ?\n, {row, _, state} ->
            {row + 1, 0, state}

          letter, {row, col, state} ->
            {row, col + 1, upsert_mapset(state, {row, col}, <<letter::utf8>>)}
        end
      )

    # -1 cause \n affect row length
    {rows, cols - 1, map |> Map.delete(?.)}
  end

  defp upsert_mapset(state, {row, col}, letter) do
    case Map.get(state, letter) do
      nil -> Map.put(state, letter, MapSet.new() |> MapSet.put({row, col}))
      ms -> %{state | letter => MapSet.put(ms, {row, col})}
    end
  end
end
