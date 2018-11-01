/*
   DDS, a bridge double dummy solver.

   Copyright (C) 2006-2014 by Bo Haglund /
   2014-2016 by Bo Haglund & Soren Hein.

   See LICENSE and README.
*/


#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "../include/dll.h"

unsigned char dcardSuit[5] = { 'S', 'H', 'D', 'C', 'N' };


void PrintTable(ddTableResults * table)
{
  
  printf("%5d;%5d;%5d;%5d;",
         table->resTable[4][0],
         table->resTable[4][2],
         table->resTable[4][1],
         table->resTable[4][3]);

  for (int suit = 0; suit < DDS_SUITS; suit++)
  {
    printf("%5d;%5d;%5d;%5d;",
           table->resTable[suit][0],
           table->resTable[suit][2],
           table->resTable[suit][1],
           table->resTable[suit][3]);
  }
}

void PrintPar(parResults * par)
{
  printf("%s;%s;%s;%s",par->parScore[0],par->parScore[1],par->parContractsString[0],par->parContractsString[1]);
}

int main(int argc, char *argv[])
{
  char svul[1]  ;
  int iVul=0;
  int res;
  char line[80];
  char pbnc[80]="";
  parResults pres;
  dealPBN dlPBN;
  ddTableDealPBN tableDealPBN;
  ddTableResults table;

  if (argc>=2){
    sprintf(pbnc,"%s",argv[1]);
  }
  if (argc>=3){
    sprintf(svul,"%s",argv[2]);
    iVul=atoi(svul);
  }

  SetMaxThreads(0);
  strcpy(tableDealPBN.cards, pbnc);
  res=CalcDDtablePBN(tableDealPBN, &table);
  if (res != RETURN_NO_FAULT)
  {
    ErrorMessage(res, line);
    printf("DDS error: %s\n", line);
  } else {
     Par(&table, &pres, iVul);
     PrintPar(&pres);
     printf("\n");
     PrintTable(&table);
    }
}
