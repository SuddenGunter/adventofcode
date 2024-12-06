defmodule Input do
  @spec read(Path.t()) :: {:ok, String.t()} | {:error, String.t()}
  def read(filename) do
    case File.read(filename) do
      {:ok, contents} -> {:ok, contents}
      {:error, reason} -> {:error, "failed to read file '#{filename}': #{reason}"}
    end
  end
end
