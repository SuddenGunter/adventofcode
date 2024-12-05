defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    state = parse_input(input)

    Map.get(state, "A")
    |> Enum.flat_map(fn {row, col} ->
      [
        # m _ m
        # _ a _
        # s _ s
        %{
          "M" => [{row - 1, col - 1}, {row - 1, col + 1}],
          "S" => [{row + 1, col - 1}, {row + 1, col + 1}]
        },
        # s _ m
        # _ a _
        # s _ m
        %{
          "M" => [{row + 1, col + 1}, {row - 1, col + 1}],
          "S" => [{row + 1, col - 1}, {row - 1, col - 1}]
        },
        # s _ s
        # _ a _
        # m _ m
        %{
          "M" => [{row + 1, col - 1}, {row + 1, col + 1}],
          "S" => [{row - 1, col - 1}, {row - 1, col + 1}]
        },
        # m _ s
        # _ a _
        # m _ s
        %{
          "M" => [{row + 1, col - 1}, {row - 1, col - 1}],
          "S" => [{row + 1, col + 1}, {row - 1, col + 1}]
        }
      ]
    end)
    |> Enum.filter(fn coords ->
      Map.keys(coords)
      |> Enum.all?(fn letter ->
        all_members(state, coords, letter)
      end)
    end)
    |> Enum.count()
  end

  defp parse_input(input) do
    input
    |> String.split("", trim: true)
    |> Enum.reduce(
      {0, 0,
       %{
         "X" => MapSet.new(),
         "M" => MapSet.new(),
         "A" => MapSet.new(),
         "S" => MapSet.new()
       }},
      fn
        "\n", {x, _, state} ->
          {x + 1, 0, state}

        letter, {x, y, state} ->
          {x, y + 1, %{state | letter => MapSet.put(Map.get(state, letter), {x, y})}}
      end
    )
    |> elem(2)
  end

  defp all_members(state, coords, letter) do
    Map.get(coords, letter)
    |> Enum.all?(fn x -> MapSet.member?(Map.get(state, letter), x) end)
  end
end
