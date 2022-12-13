defmodule Comparator do
  @spec compare(any(), any()) :: boolean
  def compare(l, r) do
    case check_order(l, r) do
      :ok -> true
      :inconclusive -> true
      :wrong -> false
    end
  end

  @spec check_order(any(), any()) :: :ok | :wrong | :inconclusive
  defp check_order(l, r) do
    case {l, r} do
      {a, b} when is_integer(a) and is_integer(b) -> compare_ints(a, b)
      {a, b} when is_list(a) and is_list(b) -> compare_lists(a, b)
      {a, b} when is_integer(a) and is_list(b) -> check_order([a], b)
      {a, b} when is_list(a) and is_integer(b) -> check_order(a, [b])
    end
  end

  @spec compare_ints(integer(), integer()) :: :ok | :wrong | :inconclusive
  defp compare_ints(a, b) do
    cond do
      a < b -> :ok
      a > b -> :wrong
      a == b -> :inconclusive
    end
  end

  @spec compare_lists([], []) :: :ok | :wrong | :inconclusive
  defp compare_lists(a, b) do
    cond do
      a == [] and b == [] -> :inconclusive
      a == [] -> :ok
      b == [] -> :wrong
      true -> compare_lists_by_element(a, b)
    end
  end

  @spec compare_lists_by_element([], []) :: :ok | :wrong | :inconclusive
  defp compare_lists_by_element(a, b) do
    first_a = Enum.at(a, 0)
    first_b = Enum.at(b, 0)
    order = check_order(first_a, first_b)

    if order != :inconclusive do
      order
    else
      compare_lists(Enum.drop(a, 1), Enum.drop(b, 1))
    end
  end
end
