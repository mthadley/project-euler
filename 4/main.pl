#!/usr/bin/env perl6

sub is_pal(Str $num) { $num eq $num.flip }

my $largest = 0;

for 100..999 -> $x {
  for 100..999 -> $y {
    my $prod = $x * $y;

    $largest = $prod if is_pal(~$prod) && $prod > $largest;
  }
}

say $largest;
