defmodule Task2 do
  @spec solution([%Reindeer{}]) :: integer()
  def solution(deers) do
    solution(deers, 2503)
  end

  @spec solution([%Reindeer{}], integer()) :: integer()
  def solution(deers, seconds) do
    score(deers, seconds)
    |> Enum.map(fn x -> x.score end)
    |> Enum.max()
  end

  defp score(deers, seconds) do
    range = 0..seconds

    deer_scores =
      Enum.map(deers, fn x ->
        %Racer{
          score: 0,
          distance: 0,
          runTime: x.runTime,
          restTime: x.restTime,
          action: :running,
          canDoAction: x.runTime,
          speed: x.speed
        }
      end)

    Enum.reduce(range, deer_scores, fn _, acc ->
      after_action =
        Enum.map(
          acc,
          fn racer ->
            actionable_racer = get_action(racer)

            case actionable_racer.action do
              :running ->
                Map.put(
                  actionable_racer,
                  :distance,
                  actionable_racer.distance + actionable_racer.speed
                )
                |> Map.put(:canDoAction, actionable_racer.canDoAction - 1)

              :resting ->
                Map.put(actionable_racer, :canDoAction, actionable_racer.canDoAction - 1)
            end
          end
        )

      max = Enum.map(after_action, fn x -> x.distance end) |> Enum.max()

      Enum.map(after_action, fn x ->
        cond do
          x.distance >= max ->
            Map.put(x, :score, x.score + 1)

          true ->
            x
        end
      end)
    end)
  end

  defp opposite(action) do
    case action do
      :running -> :resting
      :resting -> :running
    end
  end

  defp action_duration(racer) do
    case racer.action do
      :running -> racer.runTime
      :resting -> racer.restTime
    end
  end

  defp get_action(racer) do
    cond do
      racer.canDoAction == 0 ->
        with_action = Map.put(racer, :action, opposite(racer.action))
        Map.put(with_action, :canDoAction, action_duration(with_action))

      true ->
        racer
    end
  end
end
