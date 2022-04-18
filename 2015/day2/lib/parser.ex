defmodule Parser do
  @spec dimensions(String.t()) :: [[integer()]]
  def dimensions(contents) do
    contents
    |> String.split("\n", trim: true)
    |> Enum.map(fn x -> parseLine(x) end)
  end

  @spec parseLine(String.t()) :: [integer()]
  def parseLine(line) do
    line |> String.split("x", trim: true) |> Enum.map(fn x -> String.to_integer(x) end)
  end
end
