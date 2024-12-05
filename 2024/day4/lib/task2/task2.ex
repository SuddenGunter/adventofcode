defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(memory) do
    Regex.scan(~r/mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)/, memory)
    |> Enum.reduce(
      {:enabled, 0},
      fn
        ["do()"], {_, sum} ->
          {:enabled, sum}

        ["don't()"], {_, sum} ->
          {:disabled, sum}

        [_, _, _], {:disabled, sum} ->
          {:disabled, sum}

        [_, a, b], {:enabled, sum} ->
          {:enabled, sum + String.to_integer(a) * String.to_integer(b)}
      end
    )
    |> elem(1)
  end
end
