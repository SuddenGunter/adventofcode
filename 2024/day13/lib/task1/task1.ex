defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    input
    |> parse()
    |> Enum.map(&solve/1)
    |> Enum.map(fn {a, b} ->
      3 * a + b
    end)
    |> Enum.sum()
  end

  def parse(input) do
    input
    |> String.split("\n\n", trim: true)
    |> Enum.map(fn machine ->
      [ax, ay, bx, by, px, py] =
        Regex.scan(~r/\d+/, machine) |> Enum.map(fn [x] -> String.to_integer(x) end)

      %{:a => {ax, ay}, :b => {bx, by}, :prize => {px, py}}
    end)
  end

  def solve(%{a: {x1, x2}, b: {y1, y2}, prize: {z1, z2}}) do
    d = x1 * y2 - y1 * x2

    if d == 0 do
      {0, 0}
    else
      db = z2 * x1 - z1 * x2

      b = div(db, d)
      da = z1 * y2 - y1 * z2
      a = div(da, d)

      cond do
        x1 * a + y1 * b != z1 ->
          {0, 0}

        x2 * a + y2 * b != z2 ->
          {0, 0}

        true ->
          {a, b}
      end
    end
  end
end
