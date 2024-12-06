defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    {_, _, map} = parse_input(input)

    try_block(map)
  end

  defp try_block(map) do
    # no reason to try to block any block outside of initial path - will never be reached
    # also need to remove initial pos if it was ever visited
    initial_path =
      walk(map, MapSet.new())
      |> MapSet.to_list()
      |> Enum.map(fn {x, _} -> x end)
      |> Enum.uniq()
      |> Enum.filter(fn x -> x != map.guard end)

    initial_path
    |> Enum.map(fn {row, col} ->
      updated_map = %{
        map
        | :blocked => map.blocked |> MapSet.put({row, col}),
          :empty => map.empty |> MapSet.delete({row, col})
      }

      case walk(updated_map, MapSet.new()) do
        :loop -> 1
        _ -> 0
      end
    end)
    |> Enum.sum()
  end

  defp walk(map, visited) do
    {old_row, old_col} = map.guard

    next_pos =
      case map.guard_direction do
        :up -> {old_row - 1, old_col}
        :right -> {old_row, old_col + 1}
        :left -> {old_row, old_col - 1}
        :down -> {old_row + 1, old_col}
      end

    with true <- not MapSet.member?(visited, {next_pos, map.guard_direction}) do
      case at(map, next_pos) do
        :outside_of_map ->
          visited

        :blocked ->
          change_direction(map) |> walk(visited)

        :empty ->
          move(map, next_pos) |> walk(MapSet.put(visited, {next_pos, map.guard_direction}))
      end
    else
      false -> :loop
    end
  end

  defp at(map, pos) do
    cond do
      MapSet.member?(map.empty, pos) -> :empty
      MapSet.member?(map.blocked, pos) -> :blocked
      true -> :outside_of_map
    end
  end

  defp move(map, next_pos) do
    old_pos = map.guard
    # IO.inspect(format_map(map), label: :before)

    updmap = %{
      map
      | :guard => next_pos,
        :empty => MapSet.put(map.empty, old_pos) |> MapSet.delete(next_pos)
    }

    # IO.inspect(format_map(updmap), label: :after)

    updmap
  end

  defp format_map(map) do
    merged = (map.empty |> MapSet.to_list()) ++ (map.blocked |> MapSet.to_list())

    (merged ++ [map.guard])
    |> Enum.sort()
    |> Enum.chunk_by(fn {row, _} -> row end)
    |> Enum.map(fn row ->
      row
      |> Enum.map(fn x ->
        cond do
          MapSet.member?(map.empty, x) -> "."
          MapSet.member?(map.blocked, x) -> "#"
          true -> "O"
        end
      end)
    end)
  end

  defp change_direction(%{:guard_direction => dir} = map) do
    new_dir =
      case dir do
        :up -> :right
        :right -> :down
        :down -> :left
        :left -> :up
      end

    %{map | :guard_direction => new_dir}
  end

  defp parse_input(input) do
    input
    |> String.to_charlist()
    |> Enum.with_index()
    |> Enum.reduce(
      {0, 0,
       %{
         :blocked => MapSet.new(),
         :empty => MapSet.new(),
         :guard => {-1, -1},
         # hardcoded by default
         :guard_direction => :up
       }},
      fn
        {?\n, pos}, {row, _, acc} ->
          {row + 1, pos + 1, acc}

        {?#, pos}, {row, rowOffset, acc} ->
          {row, rowOffset, %{acc | :blocked => MapSet.put(acc.blocked, {row, pos - rowOffset})}}

        {?., pos}, {row, rowOffset, acc} ->
          {row, rowOffset, %{acc | :empty => MapSet.put(acc.empty, {row, pos - rowOffset})}}

        {?^, pos}, {row, rowOffset, acc} ->
          {row, rowOffset, %{acc | :guard => {row, pos - rowOffset}}}
      end
    )
  end
end
