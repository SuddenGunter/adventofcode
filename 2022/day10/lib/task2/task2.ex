defmodule Task2 do
  defmodule State do
    defstruct interesting_steps: Map.new(), current_value: 1

    def new() do
      %State{
        current_value: 1
      }
    end
  end

  @spec solution(Stream.t()) :: integer()
  def solution(commands) do
    commands
    |> Stream.flat_map(fn x ->
      case x do
        {:noop} -> [{:noop}]
        {:addx, arg} -> [{:noop}, {:add, arg}]
      end
    end)
    |> Stream.with_index(1)
    |> Stream.map(fn {cmd, cycle} -> {cycle, cmd} end)
    |> Enum.reduce(State.new(), &process_cycle/2)
  end

  defp process_cycle({cycle, cmd}, state) do
    pixel = draw(state.current_value, cycle)
    IO.ANSI.format([:bright, pixel], true) |> IO.write()
    :timer.sleep(15)

    %State{
      current_value: apply_cmd(state.current_value, cmd)
    }
  end

  defp draw(val, cycle) do
    pixel_pos =
      case rem(cycle, 40) do
        0 -> 40
        n -> n
      end

    pixel =
      if pixel_pos in val..(val + 2) do
        "█"
      else
        " "
      end

    newline =
      if pixel_pos == 40 do
        "\n"
      else
        ""
      end

    pixel <> newline
  end

  defp apply_cmd(val, cmd) do
    case cmd do
      {:noop} -> val
      {:add, arg} -> val + arg
    end
  end
end
