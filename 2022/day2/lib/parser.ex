defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: [{String.t(), String.t()}]
  def data(contents) do
    lines(contents)
    |> Enum.map(fn x -> parseEachLine(x) end)
  end

  @spec parseEachLine(String.t()) :: {String.t(), String.t()}
  defp parseEachLine(line) do
    [left | [right]] = String.split(line, " ")
    {left, right}
  end
end
