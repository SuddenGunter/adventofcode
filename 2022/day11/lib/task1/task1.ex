defmodule Task1 do
  defmodule State do
    defstruct monkeys: [], inspects: %{}, in_flight: %{}

    @type t :: %__MODULE__{
            monkeys: [Monkey.t()],
            inspects: %{},
            in_flight: %{}
          }

    @spec new(any) :: Task1.State.t()
    def new(monkeys) do
      inspects = monkeys |> Enum.map(fn x -> {x.id, 0} end) |> Map.new()

      %State{
        monkeys: monkeys,
        inspects: inspects,
        in_flight: %{}
      }
    end
  end

  @spec solution([Monkey.t()]) :: integer()
  def solution(monkeys) do
    simulate(monkeys, 20)
  end

  defp simulate(monkeys, rounds) do
    Range.new(1, rounds)
    |> Enum.reduce(State.new(monkeys), &round/2)
    |> Map.get(:inspects)
    |> Map.values()
    |> Enum.sort()
    |> Enum.take(2)
    |> Enum.reduce(1, fn x, acc -> x * acc end)
  end

  defp round(_, state) do
    Enum.reduce(state.monkeys, state, &turn/2)
  end

  @spec turn(Monkey.t(), State.t()) :: State.t()
  defp turn(monkey, state) do
    {with_received, cleared_in_flight} = process_in_flight(monkey, state.in_flight)

    new_in_flight =
      with_received
      |> Enum.reduce(cleared_in_flight, fn x, acc ->
        result = monkey.operation.(x)
        level = div(result, 3)
        target = monkey.test.(level)
        Map.update(acc, target, [], fn ex -> [level] ++ ex end)
      end)

    updated_monkey = %{monkey | items: []}

    %State{
      inspects: Map.update!(state.inspects, monkey.id, fn old -> old + length(with_received) end),
      monkeys:
        Enum.take(state.monkeys, monkey.id) ++
          [updated_monkey] ++ Enum.drop(state.monkeys, monkey.id + 1),
      in_flight: new_in_flight
    }
  end

  defp process_in_flight(monkey, in_flight) do
    {delivered, m} = Map.pop(in_flight, monkey.id, [])
    {monkey.items ++ delivered, m}
  end
end
