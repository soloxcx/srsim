import { event } from "@srsim/ts-types";
import React from "react";
import { LogCard } from "./LogCard";
import { FcSportsMode } from "react-icons/fc";
import { JSONTree } from "react-json-tree";
import { TurnOrder } from "./TurnOrder";

type Prop = {
  names: { [key: string]: string };
  event: {
    name: "TurnStart";
    event: event.TurnStart;
  };
};

export const TurnStart = ({ names, event }: Prop) => {
  let e = event;
  return (
    <LogCard
      className="pt-6"
      msg={`${names[e.event.active] ?? e.event.active}'s turn (current av: ${e.event.total_av.toFixed(0)})`}
      icon={<FcSportsMode />}
    >
      <TurnOrder names={names} turns={e.event.turn_order} />
    </LogCard>
  );
};
