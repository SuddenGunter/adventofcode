defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(line) do
    find_captures(line)
    |> List.flatten()
    |> Enum.map(fn x -> String.to_integer(x) end)
    |> Enum.sum()
  end

  defp find_captures(line) do
    Regex.scan(~r/-?\d+/, line) ++ ["0"]
  end
end
