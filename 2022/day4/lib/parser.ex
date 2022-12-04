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

  @spec parseEachLine(String.t()) :: {{Integer.t(), Integer.t()}, {Integer.t(), Integer.t()}}
  defp parseEachLine(line) do
    [left | [right]] = String.split(line, ",")
    {parseInteval(left), parseInteval(right)}
  end

  defp parseInteval(interval) do
    [l | [r]] = String.split(interval, "-")
    {num_l, _} = Integer.parse(l)
    {num_r, _} = Integer.parse(r)
    {num_l, num_r}
  end
end
