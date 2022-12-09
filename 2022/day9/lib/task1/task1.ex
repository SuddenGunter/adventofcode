defmodule Task1 do
  defmodule State do
    defstruct history: MapSet.new(), h_pos: {0, 0}, t_pos: {0, 0}
  end

  @spec solution([]) :: integer()
  def solution(commands) do
    commands
    |> unravel()
    |> process()
    |> MapSet.to_list()
    |> length()
  end

  # transforms "U 2, L 1" to "U U L"
  defp unravel(commands) do
    commands
    |> Enum.map(fn {x, num} ->
      Range.new(1, num)
      |> Enum.map(fn _ -> x end)
    end)
    |> List.flatten()
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

    # debug_step(h_pos, t_pos)

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

  defp debug_step(h, t) do
    IO.puts("")

    Range.new(6, 0)
    |> Enum.each(fn x ->
      draw_line(x, h, t)
    end)

    IO.puts("")
  end

  defp draw_line(x, h, t) do
    Range.new(0, 6)
    |> Enum.each(fn y ->
      case {y, x} do
        {0, 0} -> IO.write("S")
        n when n == h -> IO.write("H")
        n when n == t -> IO.write("T")
        _ -> IO.write(".")
      end
    end)

    IO.write("\n")
  end
end
