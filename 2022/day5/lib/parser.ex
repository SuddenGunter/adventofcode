defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: {[], []}
  def data(contents) do
    l = lines(contents)
    {stacks, next} = parseStacks(l)
    notEmpty = Enum.map(stacks, &removeEmpty/1)
    actions = parseActions(l, next + 1)
    {notEmpty, actions}
  end

  defp removeEmpty(stack) do
    Enum.filter(stack, fn x -> Enum.any?(x) end)
  end

  defp parseStacks(l) do
    parseNextStackRow(l, 0)
  end

  defp parseNextStackRow(l, i) do
    row =
      Enum.chunk_every(String.graphemes(Enum.at(l, i)), 4)
      |> Enum.map(fn x -> Enum.filter(x, fn el -> String.match?(el, ~r/^[a-z]$/i) end) end)

    if Enum.reduce(row, false, fn x, acc -> Enum.any?(x) || acc end) do
      {next_row, next_line} = parseNextStackRow(l, i + 1)
      {placeOnTop(row, next_row), next_line}
    else
      {row, i}
    end
  end

  defp placeOnTop(row, next_row) do
    Enum.zip(row, next_row) |> Enum.map(fn {l, r} -> [l | r] end)
  end

  defp parseActions(l, i) do
    Enum.slice(l, i..-1) |> Enum.map(&parseAction/1)
  end

  defp parseAction(l) do
    [count | [from | [to]]] =
      Regex.run(~r/^move (\d+) from (\d+) to (\d+)$/, l)
      |> Enum.drop(1)
      |> Enum.map(fn x ->
        {n, _} = Integer.parse(x)
        n
      end)

    {count, from, to}
  end
end

# Parse line Enum.chunk_every(String.graphemes(s),4) |> Enum.map(fn x -> Enum.filter(x, fn el -> String.match?(el, ~r/^[a-z]$/i) end) end)
