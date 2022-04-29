defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.downcase()
    |> String.split("\n", trim: true)
  end
end
