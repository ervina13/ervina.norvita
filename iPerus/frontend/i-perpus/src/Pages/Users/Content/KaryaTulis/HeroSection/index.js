import React, { useState } from "react";
import { Button } from "../ButtonElement";
import {
  HeroContainer,
  HeroBg,
  ImgBg,
  HeroContent,
  HeroH1,
  HeroP,
  HeroBtnWrapper,
  ArrowForward,
  ArrowRight,
  Input,
  Search,
  Btn,
} from "./HeroElements";
import IMG from "../images/img-2.jpg";

const HeroSection = () => {
  const [hover, setHover] = useState(false);

  const onHover = () => {
    setHover(!hover);
  };

  return (
    <HeroContainer>
      <HeroBg>
        <ImgBg src={IMG} />
      </HeroBg>
      <HeroContent>
        <HeroH1>Karya Tulis</HeroH1>
        {/* <HeroBtnWrapper>
          <Search>
            <Input />
            <Input />
          </Search>
        </HeroBtnWrapper> */}
        {/* <Button
          to=""
          onMouseEnter={onHover}
          onMouseLeave={onHover}
          primary="true"
          dark="true"
        >
          Cari {hover ? <ArrowForward /> : <ArrowRight />}
        </Button> */}
      </HeroContent>
    </HeroContainer>
  );
};

export default HeroSection;
