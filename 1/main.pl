#!/usr/bin/env perl6

say [+] gather { .take if $_ % 5 == 0 || $_ % 3 == 0 for 1..^1000 }

# vim: ft=perl6
