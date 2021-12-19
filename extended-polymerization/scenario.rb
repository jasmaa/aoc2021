class Scenario
  def initialize(polymer, rules)
    @polymer = polymer.chars
    @rules = rules
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
    Scenario.new polymer, rules
  end

  def step
    insertions = []
    insertion_offset = 0
    (1...@polymer.length).each do |i|
      pair = "#{@polymer[i-1]}#{@polymer[i]}"
      output = @rules[pair]
      if output
        insertions << [i + insertion_offset, output]
        insertion_offset += 1
      end
    end
    insertions.each do |insertion|
      @polymer.insert insertion[0], insertion[1]
    end
  end

  def commonality_difference
    assoc_counts = element_counts
    most_common_counts = assoc_counts.max_by{ |h| h[:count] }[:count]
    least_common_counts = assoc_counts.min_by{ |h| h[:count] }[:count]
    most_common_counts - least_common_counts
  end

  private

  def element_counts
    counts = {}
    @polymer.each do |element|
      if counts[element]
        counts[element] += 1
      else
        counts[element] = 1
      end
    end
    counts.map{ |k, v| {element: k, count: v} }.reduce([], :<<)
  end
end