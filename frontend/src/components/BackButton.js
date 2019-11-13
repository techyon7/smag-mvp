import React, { Component, useState } from "react";
import { withRouter, history } from "react-router";
// import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import backicon from "./img/chevron-left-solid.svg";

function BackButton() {
  return (
    <div className="BackButton">
      <a href="/">
        <img className="BackButtonImg" src={backicon} width="24px" />
      </a>
    </div>
  );
}

export default BackButton;
