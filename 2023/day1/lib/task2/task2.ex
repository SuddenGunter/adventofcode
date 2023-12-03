defmodule Task2 do
  @spec solution([String.t()]) :: integer()
  def solution(lines) do
    Enum.map(lines, &replace_words_with_digits/1) |> Task1.solution()
  end

  @spec replace_words_with_digits(String.t()) :: String.t()
  def replace_words_with_digits(l) do
    num = first_num(l)

    case num do
      {_, word, digit} ->
        String.replace(l, word, digit, global: false) |> replace_words_with_digits()

      :error ->
        l
    end
  end

  @spec first_num(String.t()) :: {String.t(), String.t()} | {:error}
  def first_num(l) do
    word_digits = %{
      "one" => "1e",
      "two" => "2o",
      "three" => "3e",
      "four" => "4",
      "five" => "5e",
      "six" => "6",
      "seven" => "7n",
      "eight" => "8t",
      "nine" => "9e"
    }

    Enum.map(word_digits, fn {k, v} -> {pos(l, k), k, v} end)
    |> Enum.filter(fn {pos, _, _} -> pos >= 0 end)
    |> Enum.min_by(fn {pos, _, _} -> pos end, fn -> :error end)
  end

  @spec pos(String.t(), String.t()) :: integer()
  def pos(l, k) do
    match = :binary.match(l, k)

    case match do
      :nomatch -> -1
      {idx, _} -> idx
    end
  end
end
