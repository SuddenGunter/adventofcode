defmodule Task1 do
  use Bitwise
  use Memoize

  @spec solution(%{String.t() => %Signal{}}) :: integer()
  def solution(signals) do
    signal = signals["a"]
    eval(signal, signals)
  end

  defmemo eval(signal, _) when is_integer(signal.left) do
    signal.left
  end

  defmemo eval(signal, signals) do
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
