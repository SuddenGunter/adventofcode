defmodule Parser do
  @spec line(String.t()) :: String.t()
  def line(contents) do
    contents
    |> String.downcase()
    |> String.trim()
  end
end
