defmodule Task1 do
  @spec solution([[integer()]]) :: integer()
  def solution(reports) do
    reports
    |> Enum.map(&safe_level?/1)
    |> Enum.count(fn x -> x end)
  end

  @spec safe_level?([integer()]) :: boolean()
  def safe_level?([h | [next | _]] = data) do
    cond do
      h > next -> safe_level?(data, :dec)
      h < next -> safe_level?(data, :inc)
      h == next -> false
    end
  end

  @spec safe_level?([integer()], atom()) :: boolean()
  def safe_level?([h | [next | _] = tail], :inc) do
    cond do
      h < next and safe_distance?(h, next) -> safe_level?(tail, :inc)
      true -> false
    end
  end

  @spec safe_level?([integer()], atom()) :: boolean()
  def safe_level?([h | [next | _] = tail], :dec) do
    cond do
      h > next and safe_distance?(h, next) -> safe_level?(tail, :dec)
      true -> false
    end
  end

  @spec safe_level?([integer()], any()) :: boolean()
  def safe_level?(data, _) when length(data) == 1 do
    true
  end

  @spec safe_level?([], any()) :: boolean()
  def safe_level?([], _) do
    true
  end

  def safe_distance?(x, y) do
    distance = abs(x - y)

    cond do
      distance < 1 -> false
      distance > 3 -> false
      true -> true
    end
  end
end
