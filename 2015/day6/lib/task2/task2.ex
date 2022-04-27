defmodule Task2 do
  @spec solution([%Command{}]) :: integer()
  def solution(commands) do
    commands
    |> Enum.reduce(%{}, fn x, acc -> simulateStep(x, acc) end)
    |> Map.values()
    |> Enum.sum()
  end

  defp simulateStep(step, acc) do
    pos = getAllPositions(step)

    pos
    |> Enum.reduce(acc, fn x, acc ->
      Map.put(
        acc,
        x,
        getValue(Map.get(acc, x, 0), step.action)
      )
    end)
  end

  defp getValue(oldVal, action) when oldVal <= 0 do
    case action do
      :toggle -> oldVal + 2
      :turn_on -> oldVal + 1
      :turn_off -> 0
    end
  end

  defp getValue(oldVal, action) when oldVal > 0 do
    case action do
      :toggle -> oldVal + 2
      :turn_on -> oldVal + 1
      :turn_off -> oldVal - 1
    end
  end

  defp getAllPositions(step) do
    step.from.x..step.to.x
    |> Stream.flat_map(fn xPos ->
      step.from.y..step.to.y |> Stream.map(fn yPos -> %Position{x: xPos, y: yPos} end)
    end)
  end
end
