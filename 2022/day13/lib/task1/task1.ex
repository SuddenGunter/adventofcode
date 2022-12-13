defmodule Task1 do
  @spec solution([{[], []}]) :: integer()
  def solution(pairs) do
    Enum.map(pairs, &check_order/1)
    |> Enum.with_index()
    |> Enum.filter(fn {x, _} -> x == :ok end)
    |> Enum.map(fn {_, i} -> i + 1 end)
    |> Enum.sum()
  end

  @spec solution({any(), any()}) :: :atom
  defp check_order(pair) do
    case pair do
      {a, b} when is_integer(a) and is_integer(b) -> compare_ints(a, b)
      {a, b} when is_list(a) and is_list(b) -> compare_lists(a, b)
      {a, b} when is_integer(a) and is_list(b) -> check_order({[a], b})
      {a, b} when is_list(a) and is_integer(b) -> check_order({a, [b]})
    end
  end

  defp compare_ints(a, b) do
    cond do
      a < b -> :ok
      a > b -> :wrong
      a == b -> :inconclusive
    end
  end

  defp compare_lists(a, b) do
    cond do
      a == [] and b == [] -> :inconclusive
      a == [] -> :ok
      b == [] -> :wrong
      true -> compare_lists_by_element(a, b)
    end
  end

  defp compare_lists_by_element(a, b) do
    first_a = Enum.at(a, 0)
    first_b = Enum.at(b, 0)
    order = check_order({first_a, first_b})

    if order != :inconclusive do
      order
    else
      compare_lists(Enum.drop(a, 1), Enum.drop(b, 1))
    end
  end
end
