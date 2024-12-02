defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: [[integer()]]
  def data(contents) do
    lines(contents)
    |> Enum.map(fn x -> String.split(x, " ", trim: true) end)
    |> Enum.map(fn x ->
      Enum.map(x, &String.to_integer/1)
    end)
  end
end
