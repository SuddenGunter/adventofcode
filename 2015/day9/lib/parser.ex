defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec distances(String.t()) :: %{}
  def distances(contents) do
    contents
    |> lines
    |> Enum.reduce(%{}, fn x, acc ->
      {from, to, weight} = parse_line(x)
      Map.put(acc, {from, to}, weight) |> Map.put({to, from}, weight)
    end)
  end

  defp parse_line(line) do
    [path | [weight]] = String.split(line, "=", trim: true)
    [from | [to]] = path |> String.split("to", trim: true) |> Enum.map(fn x -> String.trim(x) end)
    {from, to, weight |> String.trim() |> String.to_integer()}
  end
end
