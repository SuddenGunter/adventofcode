defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(inputs) do
    String.codepoints(inputs)
    |> Enum.with_index()
    |> Enum.reduce_while(0, fn {ch, index}, acc ->
      case ch do
        "(" ->
          {:cont, acc + 1}

        ")" ->
          if acc - 1 < 0 do
            {:halt, index + 1}
          else
            {:cont, acc - 1}
          end
      end
    end)
  end
end
