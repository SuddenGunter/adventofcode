defmodule Task1 do
  defmodule State do
    defstruct interesting_steps: Map.new(), current_value: 1

    def new() do
      %State{
        interesting_steps: %{
          20 => nil,
          60 => nil,
          100 => nil,
          140 => nil,
          180 => nil,
          220 => nil
        },
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
    |> Map.get(:interesting_steps)
    |> Map.values()
    |> Enum.sum()
  end

  defp process_cycle({cycle, cmd}, state) do
    %State{
      current_value: apply_cmd(state.current_value, cmd),
      interesting_steps: Map.replace(state.interesting_steps, cycle, state.current_value * cycle)
    }
  end

  defp apply_cmd(val, cmd) do
    case cmd do
      {:noop} -> val
      {:add, arg} -> val + arg
    end
  end
end
