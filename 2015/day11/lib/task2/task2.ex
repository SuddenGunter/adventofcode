defmodule Task2 do
  @spec solution(String.t()) :: String.t()
  def solution(line) do
    Task1.solution(line) |> Task1.solution()
  end
end
