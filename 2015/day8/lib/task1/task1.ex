defmodule Task1 do
  @spec solution([String.t()]) :: integer()
  def solution(lines) do
    lines
    |> Enum.reduce(0, fn x, acc ->
      acc + String.length(x) - valueLength(x)
    end)
  end

  defp valueLength(str) do
    String.graphemes(str)
    |> Enum.slice(1..(String.length(str) - 2))
    |> Enum.reduce(%{total: 0, state: :char}, fn x, acc ->
      cond do
        acc.state == :char && x == "\\" -> %{total: acc.total, state: :prevIsSlash}
        acc.state == :char && x != "\\" -> %{total: acc.total + 1, state: :char}
        acc.state == :prevIsSlash && x == "x" -> %{total: acc.total + 1, state: :skip2More}
        acc.state == :prevIsSlash && x != "x" -> %{total: acc.total + 1, state: :char}
        acc.state == :skip2More -> %{total: acc.total, state: :skip1More}
        acc.state == :skip1More -> %{total: acc.total, state: :char}
      end
    end)
    |> Map.get(:total, 0)
  end
end
