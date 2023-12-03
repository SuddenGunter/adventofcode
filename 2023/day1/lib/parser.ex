defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: String.t()
  def data(contents) do
    lines(contents)
    |> Enum.at(0)
  end
end
