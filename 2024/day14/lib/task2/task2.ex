defmodule Task2 do
  @spec solution(String.t()) :: integer()
  def solution(input) do
    size = {103, 101}

    input
    |> parse()
    |> simulate(50000, size, 0)
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

  def simulate(_, 0, _, _) do
    0
  end

  def simulate(_, _, _, 3) do
    0
  end

  def simulate(robots, steps, {map_rows, map_cols}, renders_before) do
    updated =
      robots
      |> Enum.map(fn robot ->
        {r, c} = robot.pos
        {vr, vc} = robot.vel

        %{pos: {limit(r + vr, map_rows), limit(c + vc, map_cols)}, vel: {vr, vc}}
      end)

    renders = render(updated, steps, renders_before)

    simulate(updated, steps - 1, {map_rows, map_cols}, renders)
  end

  def render(robots, step, renders_before) do
    pos =
      robots
      |> Enum.map(fn %{pos: pos} -> pos end)
      |> Enum.uniq()

    if length(pos) == length(robots) do
      img =
        Image.new!(101 * 10, 103 * 10)
        |> Image.Draw.flood!(0, 0, color: [0, 0, 0])

      last_img =
        Enum.reduce(robots, img, fn r, img ->
          {r, c} = r.pos
          Image.Draw.rect!(img, c * 10, r * 10, 10, 10, color: [0, 255, 0])
          # Image.Draw.point!(img, c, r, color: [0, 255, 0])
        end)

      Image.write!(last_img, "images/" <> Integer.to_string(50000 - step + 1) <> ".png")
      renders_before + 1
    else
      renders_before
    end
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
end
