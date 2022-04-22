defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(base) do
    mine(base, 0)
  end

  @spec mine(String.t(), integer()) :: integer()
  defp mine(base, num) do
    case hash(base, num) do
      "00000" -> num
      _ -> mine(base, num + 1)
    end
  end

  defp hash(base, num) do
    str = base <> Integer.to_string(num)
    :crypto.hash(:md5, str) |> Base.encode16() |> String.slice(0, 5)
  end
end
