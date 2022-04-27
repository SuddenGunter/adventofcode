defmodule Task1 do
  @spec solution([%Command{}]) :: integer()
  def solution(commands) do
    commands
    |> Enum.reduce(%{}, fn x, acc -> simulateStep(x, acc) end)
    |> Map.filter(fn {_, v} -> v end)
    |> Enum.count()
  end

  defp simulateStep(step, acc) do
    pos = getAllPositions(step)

    pos
    |> Enum.reduce(acc, fn x, acc ->
      Map.put(
        acc,
        x,
        case step.action do
          :toggle -> !Map.get(acc, x, false)
          :turn_on -> true
          :turn_off -> false
        end
      )
    end)
  end

  defp getAllPositions(step) do
    step.from.x..step.to.x
    |> Stream.flat_map(fn xPos ->
      step.from.y..step.to.y |> Stream.map(fn yPos -> %Position{x: xPos, y: yPos} end)
    end)
  end
end
