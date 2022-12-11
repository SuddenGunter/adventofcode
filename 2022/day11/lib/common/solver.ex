defmodule Solver do
  defmodule State do
    @enforce_keys [:monkeys, :inspects, :in_flight, :manage]
    defstruct monkeys: [], inspects: %{}, in_flight: %{}, manage: nil

    @type t :: %__MODULE__{
            monkeys: [Monkey.t()],
            inspects: %{},
            in_flight: %{},
            manage: (integer() -> integer())
          }

    @spec new([Monkey.t()], (integer() -> integer())) :: Solver.State.t()
    def new(monkeys, manage) do
      inspects = monkeys |> Enum.map(fn x -> {x.id, 0} end) |> Map.new()

      %State{
        monkeys: monkeys,
        inspects: inspects,
        in_flight: %{},
        manage: manage
      }
    end
  end

  @spec solution([Monkey.t()], integer(), (integer() -> integer())) :: integer()
  def solution(monkeys, rounds, manage) do
    simulate(monkeys, rounds, manage)
  end

  defp simulate(monkeys, rounds, manage) do
    Range.new(1, rounds)
    |> Enum.reduce(State.new(monkeys, manage), &round/2)
    |> Map.get(:inspects)
    |> Map.values()
    |> Enum.sort()
    |> Enum.take(-2)
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
        level = state.manage.(result)
        target = monkey.test.(level)
        Map.update(acc, target, [level], fn ex -> [level] ++ ex end)
      end)

    updated_monkey = %{monkey | items: []}

    %State{
      inspects: Map.update!(state.inspects, monkey.id, fn old -> old + length(with_received) end),
      monkeys:
        Enum.take(state.monkeys, monkey.id) ++
          [updated_monkey] ++ Enum.drop(state.monkeys, monkey.id + 1),
      in_flight: new_in_flight,
      manage: state.manage
    }
  end

  defp process_in_flight(monkey, in_flight) do
    {delivered, m} = Map.pop(in_flight, monkey.id, [])

    {monkey.items ++ delivered, m}
  end
end
