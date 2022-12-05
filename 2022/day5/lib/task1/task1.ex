defmodule Task1 do
  @spec solution({any(), any()}) :: String.t()
  def solution({stacks, actions}) do
    executeAction(stacks, fixIndexes(actions))
    |> Enum.map(fn x -> Enum.at(x, 0, "") end)
    |> Enum.join()
  end

  defp fixIndexes(actions) do
    Enum.map(actions, fn {x, from, to} -> {x, from - 1, to - 1} end)
  end

  defp executeAction(stacks, actions) when length(actions) == 0 do
    stacks
  end

  defp executeAction(stacks, actions) do
    {count, from, to} = Enum.at(actions, 0)

    updated_stacks = move(stacks, count, from, to)

    executeAction(updated_stacks, Enum.drop(actions, 1))
  end

  defp move(stacks, count, from, to) do
    from_stack = Enum.at(stacks, from, [])
    to_stack = Enum.at(stacks, to, [])

    {updated_from, updated_to} = moveCrate(from_stack, to_stack, count)

    replace(stacks, updated_from, from) |> replace(updated_to, to)
  end

  defp replace(stacks, stack, pos) do
    Enum.with_index(stacks)
    |> Enum.map(fn {x, ix} -> if ix == pos, do: stack, else: x end)
  end

  defp moveCrate(from_stack, to_stack, n) when n === 0 do
    {from_stack, to_stack}
  end

  defp moveCrate(from_stack, to_stack, n) do
    [h | updated_from] = from_stack

    updated_to = [h | to_stack]
    moveCrate(updated_from, updated_to, n - 1)
  end
end
