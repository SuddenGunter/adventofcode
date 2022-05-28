defmodule Parser do
  def lines(contents) do
    contents
    |> String.downcase()
    |> String.split("\n", trim: true)
  end

  @spec lines(String.t()) :: %{}
  def data(contents) do
    contents
    |> lines
    |> Enum.map(fn x -> String.trim(x, ".") end)
    |> Enum.map(fn x -> String.split(x, " ") end)
    |> Enum.map(fn x ->
      val =
        case Enum.at(x, 2) do
          "gain" -> Enum.at(x, 3) |> String.to_integer()
          "lose" -> -(Enum.at(x, 3) |> String.to_integer())
          _ -> throw("unexpected verb")
        end

      %{Name: Enum.at(x, 0), Val: val, NameReason: Enum.at(x, length(x) - 1)}
    end)
    |> Enum.reduce(%{}, fn x, acc ->
      IO.inspect(acc, label: "acc")

      Map.get_and_update(
        acc,
        x[:Name],
        fn old ->
          val =
            case old do
              nil -> %{}
              p -> p.costs
            end

          {old, %Person{costs: Map.put(val, x[:NameReason], x[:Val])}}
        end
      )
      |> elem(1)
    end)
    |> IO.inspect()
  end
end
