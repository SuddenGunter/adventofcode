defmodule Task1 do
  @spec solution([%Reindeer{}]) :: integer()
  def solution(deers) do
    solution(deers, 2503)
  end

  @spec solution([%Reindeer{}], integer()) :: integer()
  def solution(deers, seconds) do
    deers
    |> Enum.map(fn x -> getPosAt(x, seconds) end)
    |> Enum.max()
  end

  # defp getPosAt(deer, second) do
  #   iterations = div(second, deer.restTime - deer.runTime)
  #   deer.speed * iterations * deer.runTime
  # end

  defp getPosAt(deer, lastSecond) do
    getPosAt(deer, 0, :running, 0, lastSecond)
  end

  defp getPosAt(deer, currentSecond, state, distance, lastSecond) do
    cond do
      currentSecond >= lastSecond ->
        distance

      true ->
        case state do
          :running ->
            runTime =
              cond do
                currentSecond + deer.runTime <= lastSecond -> deer.runTime
                true -> lastSecond - currentSecond
              end

            getPosAt(
              deer,
              currentSecond + runTime,
              :resting,
              distance + deer.speed * runTime,
              lastSecond
            )

          :resting ->
            runTime =
              cond do
                currentSecond + deer.restTime <= lastSecond -> deer.restTime
                true -> lastSecond - currentSecond
              end

            getPosAt(
              deer,
              currentSecond + runTime,
              :running,
              distance,
              lastSecond
            )
        end
    end
  end
end
