defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    state = parse_input(input)

    Map.get(state, "X")
    |> Enum.flat_map(fn {row, col} ->
      [
        # horisontal word
        %{
          "M" => {row, col + 1},
          "A" => {row, col + 2},
          "S" => {row, col + 3}
        },
        # horisontal word (reversed)
        %{
          "M" => {row, col - 1},
          "A" => {row, col - 2},
          "S" => {row, col - 3}
        },
        # vertical word
        %{
          "M" => {row + 1, col},
          "A" => {row + 2, col},
          "S" => {row + 3, col}
        },
        # vertical word (reversed)
        %{
          "M" => {row - 1, col},
          "A" => {row - 2, col},
          "S" => {row - 3, col}
        },
        # diag (bottom, right)
        %{
          "M" => {row + 1, col + 1},
          "A" => {row + 2, col + 2},
          "S" => {row + 3, col + 3}
        },
        # diag (bottom, left)
        %{
          "M" => {row + 1, col - 1},
          "A" => {row + 2, col - 2},
          "S" => {row + 3, col - 3}
        },
        # diag (top, left)
        %{
          "M" => {row - 1, col - 1},
          "A" => {row - 2, col - 2},
          "S" => {row - 3, col - 3}
        },
        # diag (top, right)
        %{
          "M" => {row - 1, col + 1},
          "A" => {row - 2, col + 2},
          "S" => {row - 3, col + 3}
        }
      ]
    end)
    |> Enum.filter(fn coords ->
      Map.keys(coords)
      |> Enum.all?(fn letter ->
        MapSet.member?(Map.get(state, letter), Map.get(coords, letter))
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
end
