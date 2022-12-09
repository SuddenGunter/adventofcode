defmodule Task1 do
  defmodule State do
    defstruct history: MapSet.new(), h_pos: {0, 0}, t_pos: {0, 0}
  end

  @spec solution([]) :: integer()
  def solution(commands) do
    m = commands
    |> process()
      #|> IO.inspect()
    debug_map(m)
    m |> MapSet.to_list() |> length()
  end

  defp debug_map(m) do
    IO.puts("")

    Range.new(16, -16)
    |> Enum.each(fn x ->
      draw_line(x, m)
    end)

    IO.puts("")
  end

  defp draw_line(x, m) do
    Range.new(-16, 16)
    |> Enum.each(fn y ->
      found = MapSet.member?(m, {y,x})
      case {y, x} do
        {0, 0} -> IO.write("S")
        _ when found ->  IO.write("#")
        _ -> IO.write(".")
      end
    end)

    IO.write("\n")
  end


  defp process(commands) do
    commands
    |> Enum.reduce(%State{}, &step/2)
    |> Map.get(:history)
  end

  @spec step(atom(), %State{}) :: %State{}
  defp step(cmd, acc) do
    h_pos = next_h_pos(acc.h_pos, cmd)
    t_pos = next_t_pos(acc.t_pos, h_pos)

    %State{
      history: MapSet.put(acc.history, t_pos),
      h_pos: h_pos,
      t_pos: t_pos
    }
  end

  @spec next_h_pos({integer(), integer()}, atom()) :: {integer(), integer()}
  defp next_h_pos({x, y}, cmd) do
    case cmd do
      :U -> {x, y + 1}
      :R -> {x + 1, y}
      :L -> {x - 1, y}
      :D -> {x, y - 1}
      :UR -> {x + 1, y + 1}
      :UL -> {x - 1, y + 1}
      :DR -> {x + 1, y - 1}
      :DL -> {x - 1, y - 1}
    end
  end

  @spec next_t_pos({integer(), integer()}, {integer(), integer()}) :: {integer(), integer()}
  defp next_t_pos({tx, ty}, {hx, hy}) do
    y_diff = hy - ty
    x_diff = hx - tx

    if abs(y_diff) >= 2 || abs(x_diff) >= 2 do
      {tx + infer_steps(x_diff), ty + infer_steps(y_diff)}
    else
      {tx, ty}
    end
  end

  defp infer_steps(num) do
    case num do
      n when n > 0 -> 1
      n when n < 0 -> -1
      _ -> 0
    end
  end
end
