defmodule Parser do
  @spec lines(String.t()) :: [String.t()]
  def lines(contents) do
    contents
    |> String.split("\n", trim: true)
  end

  @spec games([String.t()]) :: [map()]
  def games(lines) do
    Enum.map(lines, &parse_game/1)
  end

  @spec parse_game(String.t()) :: map()
  def parse_game(l) do
    [game| [subsets |_]] = String.split(l, ":")
    result = %{game_number: parse_game_number(game)}
    parse_subsets(result, subsets)
  end

  @spec parse_game_number(String.t()) :: integer()
  def parse_game_number(l) do
    no_prefix = String.slice(l, 5..-1) # removes 'Game '
    Integer.parse(no_prefix)
  end

  @spec parse_subsets(map(), String.t()) :: map()
  def parse_subsets(result, l) do
    [current|rest]= String.split(l, ":")

    with_current_subset = Map.put(result, parse_subset(current))

    case rest do
     [] -> result
      _ -> parse_subsets(with_current_subset, Enum.join(rest))
    end
  end

  def parse_subset(subset) def
    
  do
end
