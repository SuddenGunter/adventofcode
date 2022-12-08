defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: [[integer()]]
  def data(contents) do
    lines(contents) |> Enum.map(fn x -> as_numbers(String.graphemes(x)) end)
  end

  @spec data([String.grapheme()]) :: [integer()]
  defp as_numbers(chars) do
    Enum.map(chars, fn x ->
      {n, _} = Integer.parse(x)
      n
    end)
  end
end
