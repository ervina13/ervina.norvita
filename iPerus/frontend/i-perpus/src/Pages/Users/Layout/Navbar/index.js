import React, { useEffect, useState } from "react";
import { FaBars } from "react-icons/fa";
import { IconContext } from "react-icons/lib";
import {
  Nav,
  NavbarContainer,
  NavLogo,
  MobileIcon,
  NavMenu,
  NavItem,
  NavLinks,
  NavBtn,
  NavBtnLink,
  Item,
} from "./NavbarElements";
import { Button } from "styled-button-component";
import {
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownDivider,
} from "styled-dropdown-component";

const Navbar = ({ toggle }) => {
  const [scrollNav, setScrollNav] = useState(false);
  const [hidden, setHidden] = useState(true);

  const changeNav = () => {
    if (window.scrollY >= 80) {
      setScrollNav(true);
    } else {
      setScrollNav(false);
    }
  };

  useEffect(() => {
    window.addEventListener("scroll", changeNav);
  }, []);

  return (
    <>
      <IconContext.Provider value={{ color: "#fff" }}>
        <Nav scrollNav={scrollNav}>
          <NavbarContainer>
            <NavLogo to="/home">iPerpus</NavLogo>
            <MobileIcon onClick={toggle}>
              <FaBars />
            </MobileIcon>
            <NavMenu>
              <NavItem>
                <NavLinks to="/home">Home</NavLinks>
              </NavItem>
              <NavItem>
                <NavLinks to="/karya-tulis">Karya Tulis</NavLinks>
              </NavItem>
              <NavItem>
                <NavLinks to="/contact">Contact</NavLinks>
              </NavItem>
            </NavMenu>
            <NavBtn>
              {/* <NavBtnLink to="/">Sign In</NavBtnLink> */}
              <Dropdown>
                <NavBtnLink dropdownToggle onClick={() => setHidden(!hidden)}>
                  Akun
                </NavBtnLink>
                <DropdownMenu hidden={hidden} toggle={() => setHidden(!hidden)}>
                  <DropdownItem>
                    <Item to="/profile">Profile</Item>
                  </DropdownItem>
                  <DropdownDivider />
                  <DropdownItem>
                    <Item to="/signin">Logout</Item>
                  </DropdownItem>
                </DropdownMenu>
              </Dropdown>
            </NavBtn>
          </NavbarContainer>
        </Nav>
      </IconContext.Provider>
    </>
  );
};

export default Navbar;
