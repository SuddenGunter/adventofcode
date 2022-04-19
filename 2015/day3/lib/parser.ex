defmodule Parser do
  @spec firstLine(String.t()) :: String.t()
  def firstLine(contents) do
    contents |> String.split("\n", trim: true) |> Enum.at(0)
  end
end
