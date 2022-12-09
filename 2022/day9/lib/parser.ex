defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: [[integer()]]
  def data(contents) do
    lines(contents) |> Enum.map(fn x -> parse_line(x) end)
  end

  defp parse_line(line) do
    [l | [r | _]] = String.split(line, " ")
    {steps, _} = Integer.parse(r)
    {String.to_atom(l), steps}
  end
end
