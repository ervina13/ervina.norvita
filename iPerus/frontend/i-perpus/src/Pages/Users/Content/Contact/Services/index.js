import React from "react";
import Icon1 from "../images/svg-1.svg";
import Icon2 from "../images/svg-4.svg";
import Icon3 from "../images/svg-2.svg";
import {
  ServicesContainer,
  ServicesH1,
  ServicesWrapper,
  ServicesCard,
  ServicesIcon,
  ServicesH2,
  ServicesP,
} from "./ServicesElements";

const Services = () => {
  return (
    <ServicesContainer id="services">
      <ServicesWrapper>
        <ServicesCard>
          <ServicesIcon src={Icon1} />
          <ServicesH2>ALAMAT</ServicesH2>
          <ServicesP>Indonesia</ServicesP>
        </ServicesCard>
        <ServicesCard>
          <ServicesIcon src={Icon2} />
          <ServicesH2>CALL NUMBER</ServicesH2>
          <ServicesP>08213191xxxx</ServicesP>
        </ServicesCard>
        <ServicesCard>
          <ServicesIcon src={Icon3} />
          <ServicesH2>EMAIL</ServicesH2>
          <ServicesP>iPerpus@gmail.com</ServicesP>
        </ServicesCard>
      </ServicesWrapper>
    </ServicesContainer>
  );
};

export default Services;
