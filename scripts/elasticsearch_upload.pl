#!/usr/bin/perl -w

use strict;

sub send_to_elastic {
   my $to_send = $_[0];
   `curl -XPOST http://localhost:9200/good/good -d '$to_send'`
}

while (<stdin>)
{
   chomp;
   my $ln = $_;
   send_to_elastic($ln);
}
