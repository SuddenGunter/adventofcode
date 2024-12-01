defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec data(String.t()) :: String.t()
  def data(contents) do
    lines(contents)
    |> Enum.map(fn x -> String.split(x, "   ", trim: true) end)
    |> Enum.map(fn [first, second] ->
      {n, _} = Integer.parse(first)
      {m, _} = Integer.parse(second)
      {n, m}
    end)
    |> Enum.reduce(
      {[], []},
      fn {l, r}, {lList, rList} ->
        {[l | lList], [r | rList]}
      end
    )
  end
end
