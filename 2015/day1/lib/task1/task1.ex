defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(inputs) do
    String.codepoints(inputs)
    |> Enum.reduce(%{:open => 0, :close => 0}, fn ch, acc ->
      case ch do
        "(" -> %{acc | :open => acc[:open] + 1}
        ")" -> %{acc | :close => acc[:close] + 1}
      end
    end)
    |> diff
  end

  @spec diff(map()) :: integer()
  defp diff(map) do
    map[:open] - map[:close]
  end
end
