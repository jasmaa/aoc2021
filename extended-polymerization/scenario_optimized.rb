class ScenarioOptimized
  def initialize(polymer, rules)
    @pairs = (1...polymer.length)
      .map{ |i| "#{polymer[i-1]}#{polymer[i]}" }
      .reduce({}){ |acc, el|
        if acc[el]
          acc[el] += 1
        else
          acc[el] = 1
        end
        acc
      }
    @rules = rules
    @first_element = polymer[0]
    @last_element = polymer[-1]
  end

  def self.parse(payload)
    lines = payload.split(/\r?\n/)
    polymer = []
    rules = {}
    lines.each_with_index do |line, i|
      if i == 0
        polymer = line
      else
        if line.length > 0
          captures = line.match(/^(\w{2})\s->\s(\w)$/).captures
          rules[captures[0]] = captures[1]
        end
      end
    end
    ScenarioOptimized.new polymer, rules
  end

  def step
    next_pairs = Hash.new(0)
    @pairs.each do |pair, count|
      if @rules[pair]
        [
          "#{pair[0]}#{@rules[pair]}",
          "#{@rules[pair]}#{pair[1]}",
        ].each do |p|
          next_pairs[p] += count
        end
      end
    end
    @pairs = next_pairs
  end

  def commonality_difference
    assoc_counts = element_counts
    most_common_counts = assoc_counts.max_by{ |h| h[:count] }[:count]
    least_common_counts = assoc_counts.min_by{ |h| h[:count] }[:count]
    most_common_counts - least_common_counts
  end

  private

  def element_counts
    counts = Hash.new(0)
    # Count elements twice by counting appearances in all pairs
    @pairs.each do |pair, count|
      counts[pair[0]] += count
      counts[pair[1]] += count
    end
    # Need to double-count end elements
    counts[@first_element] += 1
    counts[@last_element] += 1
    counts.each do |element, _|
      counts[element] /= 2
    end
    counts.map{ |k, v| {element: k, count: v} }.reduce([], :<<)
  end
end