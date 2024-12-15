defmodule Task1 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    size = {103, 101}

    input
    |> parse()
    |> simulate(100, size)
    |> safety_factor(size)
  end

  def parse(input) do
    input
    |> String.split("\n", trim: true)
    |> Enum.map(fn row ->
      [l, r] = String.split(row, " ")

      [col, row] =
        l
        |> String.replace_prefix("p=", "")
        |> String.split(",", trim: true)
        |> Enum.map(&String.to_integer/1)

      [vcols, vrows] =
        r
        |> String.replace_prefix("v=", "")
        |> String.split(",", trim: true)
        |> Enum.map(&String.to_integer/1)

      %{pos: {row, col}, vel: {vrows, vcols}}
    end)
  end

  def simulate(robots, steps, {map_rows, map_cols}) do
    robots
    |> Enum.map(fn robot ->
      {r, c} = robot.pos
      {vr, vc} = robot.vel

      {
        limit(r + rem(vr * steps, map_rows), map_rows),
        limit(c + rem(vc * steps, map_cols), map_cols)
      }
    end)
  end

  def limit(n, max) do
    cond do
      n >= max ->
        n - max

      n < 0 ->
        max + n

      true ->
        n
    end
  end

  def safety_factor(map, {map_rows, map_cols}) do
    skip_row = div(map_rows, 2)
    skip_col = div(map_cols, 2)

    grouped =
      map
      |> Enum.filter(fn
        {^skip_row, _} -> false
        {_, ^skip_col} -> false
        {_, _} -> true
      end)
      |> Enum.group_by(fn
        {row, col} ->
          cond do
            row < skip_row and col < skip_col -> 0
            row < skip_row and col > skip_col -> 1
            row > skip_row and col < skip_col -> 2
            row > skip_row and col > skip_col -> 3
          end
      end)

    Map.keys(grouped)
    |> Enum.map(fn k ->
      grouped |> Map.get(k) |> Enum.count()
    end)
    |> Enum.reduce(1, fn x, acc -> x * acc end)
  end
end
