use Bitwise

defmodule Task1 do
  @spec solution(%{String.t() => %Signal{}}) :: integer()
  def solution(signals) do
    signal = signals["a"]
    eval(signal, signals)
  end

  defp eval(signal, _) when is_integer(signal.left) do
    signal.left
  end

  defp eval(signal, signals) do
    case signal.gate do
      :value -> eval(signals[signal.left], signals)
      :not -> ~~~eval(signals[signal.left], signals)
      :and -> eval(signals[signal.left], signals) &&& eval(signals[signal.right], signals)
      :or -> eval(signals[signal.left], signals) ||| eval(signals[signal.right], signals)
      :lshift -> eval(signals[signal.left], signals) <<< eval(signals[signal.right], signals)
      :rshift -> eval(signals[signal.left], signals) >>> eval(signals[signal.right], signals)
    end
  end
end
