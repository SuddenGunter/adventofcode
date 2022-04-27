defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: [%Command{}]
  def data(contents) do
    lines(contents)
    |> Enum.map(fn x -> parseEachLine(x) end)
  end

  defp parseEachLine(line) do
    [left | [right | _]] = String.split(line, "through", trim: true)
    to = parsePair(right)

    action =
      cond do
        String.starts_with?(left, "toggle") -> :toggle
        String.starts_with?(left, "turn off") -> :turn_off
        String.starts_with?(left, "turn on") -> :turn_on
      end

    from =
      String.slice(left, String.length(Atom.to_string(action)), String.length(left))
      |> parsePair()

    %Command{
      from: from,
      to: to,
      action: action
    }
  end

  defp parsePair(pair) do
    [left | [right | _]] =
      pair
      |> String.split(",", trim: true)
      |> Enum.map(fn x -> String.trim(x) |> String.to_integer() end)

    %Position{x: left, y: right}
  end
end
