#!/usr/bin/env perl6

my @ones  = <zero one two three four five six seven eight nine>;
my @tens  = <zero ten twenty thirty forty fifty sixty seventy eighty ninety>;
my @teens = <ten eleven twelve thirteen fourteen fifteen sixteen seventeen eighteen nineteen>;

sub to-words(Int $num is copy) {
  gather given $num {
    when * == 1000 { .take for "one", "thousand" } 
    when * >= 100 {
      .take for @ones[$num div 100], "hundred";

      $num %= 100;
      if $num > 0 {
        take "and";
        proceed;
      }
    }
    when * < 10 { take @ones[$num] }
    when * < 20 { take @teens[$num % 10] }
    default {
      take @tens[$num div 10];

      my $ones = $num % 10;
      take @ones[$ones] if $ones > 0;
    }
  }
}

say [+] gather { take to-words($_).join.chars for 1..1000 };

# vim: ft=perl6
